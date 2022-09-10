import {DagDefinition, DagResults, Leaf} from './types';
import {getColor, initColors, rootColor} from "./colors";

interface ChartInfos {
    data: any[],
    links: any[],
}

export function generateChartInfo(nodes: Leaf[], fileName: string): ChartInfos {
    initColors()
    const data: any[] = [{
        name: fileName,
        value: "none",
        id: "1",
        itemStyle: {
            color: rootColor
        }
    }]
    const links: any[] = []

    let count: number = 2
    nodes.forEach(v => {
        count += 1
        count = recNodes(v, 1, data, links, count)
    })


    return {data, links}
}

export function generateTree(dag: DagResults): any {
    const tree: Leaf[] = []
    dag.dag.forEach(n => {
        tree.push(generateLeaf(n, 1, rootColor))
    })
    return tree
}

function generateLeaf(node: DagDefinition, depth: number, parentColor: string): Leaf {
    const leaf: Leaf = {
        name: node.name,
        depth: depth,
        color: getColor(depth, parentColor),
        checked: 0,
        isOpen: false,
        children: [],
        metadata: {
            def: node.def,
            file: node.file
        }
    }
    node.dependencies.forEach(n => leaf.children.push(generateLeaf(n, depth + 1, leaf.color)))

    return leaf
}

function recNodes(node: Leaf, parent: number, data: any[], links: any[], count: number): number {
    data.push({
        id: count.toString(),
        name: node.name,
        value: node.metadata,
        itemStyle: {
            color: node.color
        }
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
