import "encoding/json"
var amd ArticleMetaData
if err := datastore.Get(c, key, &amd); err != nil {
//...
ajson, err := json.Marshal(amd)
fmt.Fprintf(w, "{ \"1\": %s }", string(ajson))
