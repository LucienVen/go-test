package github

import (

)

/**
 字符串模板
 */

const temp1 = `{{.TotalCount}} issues:
    {{range .Items}}----------------------------------------
    Number: {{.Number}}
    User: {{.User.Login}}
    Title: {{.Title | printf "%.64s"}}
    Age: {{.CreatedAt | daysAgo}} days
    {{end}}`


