package generator

type (
	QueryName   string
	ResultID    string
	ResultType  int
	ResultIDGen func(prefix string, args QueryArgs, suffix string) ResultID
	QueryFn     func(QueryArgs) (interface{}, error)
	QueryArgs   map[string]interface{}
	Queries     map[QueryName]*Query
	Query       struct {
		name         QueryName
		resultType   ResultType
		fn           QueryFn
		parent       *Query
		children     Queries
		resultIDGen  ResultIDGen
		cache        map[ResultID]QueryResult
		errorHandler QueryErrorHandler
	}
	QueryErrorHandler func(QueryName, error, QueryArgs)
	QueryResult       struct {
		Type      ResultType
		QueryName QueryName
		ID        ResultID
		Result    interface{}
	}
)

var (
	NoParent   *Query
	NoChildren = Queries(map[QueryName]*Query{})
)

const (
	_ = ResultType(iota)
	Record
	Records
	RecordsTree
	Error
)

func NewQuery(name string, t ResultType, fn QueryFn, resultIDgen ResultIDGen, parent *Query, children Queries, errorHandler QueryErrorHandler) (q *Query) {
	q.name = QueryName(name)
	q.resultType = t
	q.fn = fn
	q.parent = parent
	q.children = children
	q.errorHandler = errorHandler
	q.resultIDGen = resultIDgen
	return
}

func (q *Query) Exec(args QueryArgs, handleError QueryErrorHandler) (result QueryResult, ok bool) {
	var err error
	rID := q.resultIDGen(string(q.name), args, "")
	result, ok = q.cache[rID]
	if ok {
		return
	}
	data, err := q.fn(args)
	if err != nil {
		q.errorHandler(q.name, err, args)
		handleError(q.name, err, args)
		ok = false
		return
	}
	result.Type = q.resultType
	result.QueryName = q.name
	result.ID = rID
	result.Result = data
	q.cache[rID] = result
	ok = true
	return
}

func (q *Query) Sub(name string) *Query {
	return q.children[QueryName(name)]
}
