import {DagDefinition, DagResults, EchartsLink, EchartsNode, Leaf} from './types';

interface ChartInfos {
    data: EchartsNode[],
    links: EchartsLink[],
}

export function generateTree(dag: DagResults): Leaf[] {
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
    const data: EchartsNode[] = [{
        name: "root",
        value: {
            def: "coucou",
            file: ""
        },
        id: "1"
    }]
    const links: EchartsLink[] = []

    let count: number = 2
    nodes.forEach(v => {
        count += 1
        count = recNodes(v, 1, data, links, count)
    })

    return {data, links}
}

function recNodes(node: Leaf, parent: number, data: EchartsNode[], links: EchartsLink[], count: number): number {
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
