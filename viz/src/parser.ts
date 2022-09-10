import {DagDefinition, DagResults, Leaf} from './types';

interface ChartInfos {
    data: any[],
    links: any[],
}

export function generateTree(dag: DagResults): any {
    const tree: Leaf[] = []
    dag.dag.forEach(n => {
            tree.push(generateLeaf(n))
        }
    )
    return tree
}

function generateLeaf(node: DagDefinition): Leaf {
    const leaf: Leaf = {
        name: node.name,
        checked: 0,
        isOpen: node.name === "#FS" ? true : node.name === "#WriteFile",
        children: [],
        metadata: {
            def: node.def,
            file: node.file
        }
    }
    node.dependencies.forEach(n => leaf.children.push(generateLeaf(n)))

    return leaf
}

export function generateChartInfo(nodes: Leaf[]): ChartInfos {
    const data: any[] = [{
        name: "root1",
        value: "none",
    }]
    const links: any[] = []

    let count: number = 2
    nodes.forEach(v => {
        if (v.isOpen) {
            count+=1
            count = recNodes(v, "root"+1, data, links, count)
        }
    })

    return { data, links }
}

function recNodes(node: Leaf, parent: string, data: any[], links: any[], count: number): number {
    data.push({
        name: node.name+count,
        value: node.metadata,
    })
    links.push({
        source: parent,
        target: node.name+count
    })

    const parentNB: number = count

    node.children.forEach(v => {
        count+=1
        count = recNodes(v, node.name+parentNB, data, links, count)
    })

    return count
}
