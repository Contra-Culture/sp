# SP is for "Static Pages"

`sp` consists of two major parts (utilites/build targets).

---

## Server

Server provides static pages and also is responsible for reader journey tracking.
### **Philosophy**


### **Design**

### **Usage**

### **Development**

---
## Generator

Generator is responsible for efficient static pages generation.

### **Philosophy**

### **Main concepts**

A **repository** has a view-model (VM), a number of queries that return VM-related diff data-sets: a single view-model, slice or tree of view-models that require re-rendering.

A **view model** is a data structure, that is versioned by a set of attributes.

An **attribute** is a content meta-data.

A **fragment** is a rendering function and a path-generator to store the generated fragment and a meta-data file. File store is provided as the default store. A fragment can depend on (include) other fragments.

A **page** is a final result of rendering. Pages are similar to fragments, but belong to a special registry and can be linked.

A **pages-tree** organizes pages in web-site, by structuring their urls and linking with a ``links-registry``.

A **links-registry** stores all the internal links in a tree. A set of sibling links can have its own linking

### **Recommended content concepts (not implemented by ``sp``)**

An **asset** is an aggregate of different versions of some data where each version requires its own page. Translation is a special kind of version. It is required and a first-level attribute. A versions primary key is a string generated from version attributes set by asset's keygen function. An asset can have its own kinda "index" page or a widget implementation for its versions.


### **Usage**

#### **eDSL**

To define all the data dependencies for consistent static pages generation you need to use ``eDSL``. To avoid complex configuration, your eDSL script must be provided with appropriate templates directory.

```golang
// blogging example
host := Host("My Personal Blog", func(host) {
  host.Query("settings", func(){

  }, func(dsl){

  })
  host.Query("publications", func() {

  }, func(dsl) {

  })
  host.Query("rubrics", func() {

  }, func(dsl) {

  })
  host.Query("attributeCategories", func() {

  }, func(dsl) {

  })
  host.Query("authors", func() {

  }, func(dsl) {

  })
  host.Query("pages", func() {

  }, func(dsl) {

  })



})
```
#### **Data Model and Data object**

Everything starts with data objects, that represent some portions of information. For example, a single publication and a rubric are two kinds of data objects. Each data object should be defined before any use. The data object definition must include: ``detailed``, ``preview`` and ``link`` views to be presented.

Data objects could be organized in several ordered groups. Every that group is a data object too and must include: ``detailed``, ``preview`` and ``link`` views as well. ``widget`` view is also required forgroups. Every group should include a split function to split a whole list of data objects (that could be big enough) on several pages. There is also a default split functions that split by year and by year and month.

``SP`` is responsible only for consistent static pages generation, so that its model corresponds to only that domain. First of all we have hosts, that are root directories for all the corresponding content.

The only predefined and required group kind is ``rubric``. Rubrics belong to a particular host and every asset must belong to some rubric. All the host's rubrics are available via ``data.Rubrics`` property.

Content assets represent main content on every host. Blog notes, landings, rubrics, documents, tutorials, etc. all are the examples of content assets. Every asset has many versions each of which requires a separate static page for presentation to be generated. Every asset version has unique (in the context of a single asset) set of attributes and every existing attributes set of an asset will get its own static page for presentation.

For example, you have an asset for two programming languages and with translations for two languages:

```json
{
  "programmingLanguage": "Golang",
  "language": "English",
}
```

```json
{
  "programmingLanguage": "JavaScript",
  "language": "English",
}
```

```json
{
  "programmingLanguage": "Golang",
  "language": "Русский",
}
```

```json
{
  "programmingLanguage": "JavaScript",
  "language": "Русский",
}
```

In that case, for every of this four versions of an asset would be generated a separate static page.

You can access all the assets via ``data.Assets`` property. All the attribute categories are available via ``data.AttributeCategories`` property. Every attribute category requires a so-called "partial" view for its presentation.





### **API**



### **Development**

---

