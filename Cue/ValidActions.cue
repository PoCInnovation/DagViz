#name: string
#on: [string]: {...}
#env: [string]: string | int
#id: 
    run: string

#job: {
    name: string
    "runs-on": "ubuntu-latest" 
    container?: {
        image: string
    }
    needs?: [...string]
    steps: [...{... }]
    ...
}

#jobs: [string]: #job

name: #name
env: #env
on: #on
jobs: #jobs