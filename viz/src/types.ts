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

export type Metadata = {
    def: string,
    file: string,
}

export type Leaf = {
    name: string,
    checked: 0 | 0.5 | 1
    isOpen: boolean
    children: Leaf[]
    metadata: Metadata
}
