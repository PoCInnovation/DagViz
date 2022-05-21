name: "string"
env: {
	"ENVIR":  "ENV"
	"ENVIR2": "ENV2"
	"ENVIR3": "ENV3"
}

on: {
	"push-to": [
		"string1",
		"string2",
	]
}
jobs: {
	"job": {
		"name":    "job1"
		"runs-on": "ubuntu-latest"
		"container": {
			"image": "string"
		}
		"needs": [
			"job1",
			"job2",
		]
		"steps": {
			"Steps1": {
				"name":            "string"
				"timeout-minutes": 10
				"uses":            "actions/checkout@v2"
				"with": {
					"repo": "string"
					"branch": "string"
				}
			}
		}
	}
}
