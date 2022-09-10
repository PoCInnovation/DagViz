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
    name: string
    depth: number
    color: string
    checked: number
    isOpen: boolean
    children: Leaf[]
    metadata: Metadata
}
