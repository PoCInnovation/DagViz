#Steps: {
    name?: string
    "timeout-minutes": *10 | >=0

    *{
        uses: "actions/checkout@v2"
        with?: [string]: string
    } | {
        run: [...string]
    }
    ...
}

#job: {
    name: string
    "runs-on": "ubuntu-latest"
    container?: {
        image: string
    }
    needs?: [...string]
    steps: [...#Steps]
    ...
}

name: string
env: [string]: string | int
on:  [string]: {...}
jobs: [jobName=string]: #job