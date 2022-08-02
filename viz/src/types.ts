export type DagDefinition = {
    name: string,
    file: string,
    def: string,
    dependencies: DagDefinition[],
}

export type DagResults = {
    file: string,
    dag: DagDefinition[]
}
