import { DagDefinition, DagResults } from './types';

export default function generateChart(dag: DagResults): string {
    const stack = "graph LR;"
    const fileNode = "PARENT[\"" + dag.file + "\"]"
    const table = generateTable(dag, fileNode)

    return stack + table.join("")
}

function generateTable(dag: DagResults, fileNode: string): string[] {
    const table: string[] = []
    let count: number = 0;

    dag.dag.forEach(n => {
        count += 1
        count = appendNodes(n, fileNode, table, count)
    })

    return table
}

function appendNodes(baseNode: DagDefinition, parentNode: string, table: string[], count: number): number {
    const node = generateNode(baseNode, parentNode, count)
    const parent: number = count

    table.push(node)
    baseNode.dependencies.forEach(n => {
        count += 1
        count = appendNodes(n,  parent + "[\"" + baseNode.name + "\"]", table, count)
    })
    return count
}

function generateNode(node: DagDefinition, parent: string, count: number): string {
    const link = parent + "-->"
    const style = count + "[\"" + node.name + "\"];\n"

    return link + style
}
