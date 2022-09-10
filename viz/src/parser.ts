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
        isOpen: false,
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
        name: "root",
        value: "none",
        id: "1"
    }]
    const links: any[] = []

    let count: number = 2
    nodes.forEach(v => {
        count += 1
        count = recNodes(v, 1, data, links, count)
    })

    return {data, links}
}

function recNodes(node: Leaf, parent: number, data: any[], links: any[], count: number): number {
    data.push({
        name: node.name,
        value: node.metadata,
        id: count.toString()
    })
    links.push({
        source: parent.toString(),
        target: count.toString()
    })

    const parentNB: number = count

    if (node.isOpen) {
        node.children.forEach(v => {
            count += 1
            count = recNodes(v, parentNB, data, links, count)
        })
    }
    return count
}
