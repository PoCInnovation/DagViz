import {DagDefinition, DagResults, Leaf} from './types';

export function generateChart(dag: Leaf[]): string {
    const stack = "graph LR;"
    const fileNode = "PARENT[\"TEST\"]"
    const table = generateTable(dag, fileNode)

    return stack + table.join("")
}

function generateTable(dag: Leaf[], fileNode: string): string[] {
    const table: string[] = []
    let count: number = 0;

    dag.forEach(n => {
        count += 1
        if (n.isOpen) {
            count = appendNodes(n, fileNode, table, count)
        }
    })

    return table
}

function appendNodes(baseNode: Leaf, parentNode: string, table: string[], count: number): number {
    const node = generateNode(baseNode, parentNode, count)
    const parent: number = count

    table.push(node)
    baseNode.children.forEach(n => {
        count += 1
        if (n.isOpen) {
            count = appendNodes(n, parent + "[\"" + baseNode.name + "\"]", table, count)
        }
    })
    return count
}

function generateNode(node: Leaf, parent: string, count: number): string {
    const link = parent + "-->"
    const style = count + "[\"" + node.name + "\"];\n"

    return link + style
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