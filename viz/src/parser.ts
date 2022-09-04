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
        name: "root",
        value: "none",
    }]
    const links: any[] = []

    nodes.forEach(v => {
        if (v.isOpen) {
            recNodes(v, "root", data, links)
            recursiveChart(v, data, links)
        }
    })

    return { data, links }
}

function recursiveChart(node: Leaf, data: any[], links: any[]): any {
    node.children.forEach(v => {
        recNodes(v, node.name, data, links)
    })
}

function recNodes(node: Leaf, parent: string, data: any[], links: any[]): any {
    data.push({
        name: node.name,
        value: node.metadata
    })
    links.push({
        source: parent,
        target: node.name
    })
    recursiveChart(node, data, links)
}

/*
 [
                {
                    name: "salut",
                    value: 10,
                    symbol: "circle",
                },
                {
                    name: "salut2",
                    value: 5,
                    symbol: "circle",
                }
            ],
 */

/*
: [
                {
                    source: "salut",
                    target: "salut2",
                }
            ]
 */