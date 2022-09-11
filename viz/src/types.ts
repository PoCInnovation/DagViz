export type DagDefinition = {
    name: string
    file: string
    def: string
    dependencies: DagDefinition[]
}

export type DagResults = {
    file: string
    dag: DagDefinition[]
}

export type Metadata = {
    def: string
    file: string
}

interface BaseLeaf {
    name: string
    depth: number
    color: string
    colorParams: ColorParams
    checked: number
    metadata: Metadata
}

export type Leaf =
    BaseLeaf & { isOpen: true, children: Leaf[] } |
    BaseLeaf & { isOpen: false }

export type ColorParams = {
    tintArray: string[]
    count: number
}

export type EchartsLink = {
    source: string,
    target: string,
}

export type EchartsNode = {
    id: string
    name: string
    value: string | Metadata
    itemStyle?: {
        color?: string
    }
}