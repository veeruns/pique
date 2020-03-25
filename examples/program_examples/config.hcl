log_level = "DEBUG"
name  = "simpleprogram"
check_frequency = 5

node "Action1" {
	blockname = "Init"
    plugin_name = "Secrets"
	plugin_function_name  = "getfromyck"
    plugin_function_arguments = <<EOT
    {
	"name": "veeruns",
	"hostname": ["hostname1", "hostname2", "hostname3"]
}
EOT

}

node "Action2" {
	blockname = "CheckJIRA"
    plugin_name = "jiraops"
	plugin_function_name  = "jiracheck"
    plugin_function_arguments = <<EOF
    { "name": "veeruns", "JIRA": ["ABC-123", "DEF-345","GHIK-867"]}
    EOF

}

