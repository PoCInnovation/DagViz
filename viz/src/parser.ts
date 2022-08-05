import {DagDefinition, DagResults} from './types';
import {tab} from "@testing-library/user-event/dist/tab";


type generateNodeReturn = {
    toAdd: string,
    count: number
    newNodes: DagDefinition[]
}

export default function generateChart(dag: DagResults): string {
    const style = "graph LR;"
    const fileNode = "PARENT[\"" + dag.file + "\"]"
    let stack = "graph LR;"

    let count = 0;
    function rec() : string[]{
        let table : string[] = []
        function r(baseNode: DagDefinition, parentNode: string, count: number) {
            const node = generateNode(baseNode, parentNode, count)
            let nombre: number = count
            table.push(node)
            baseNode.dependencies.map(n => {
                count += 1
                r(n,  nombre+ "[\"" + baseNode.name + "\"]", count)
            })
        }

        dag.dag.map(n => {
            console.log("hello\n")
            count += 1
            r(n, fileNode, count);
        })
        return table
    }

    const table = rec()

    console.log(table)
    return stack+ table.join("")
}

function generateNode(node: DagDefinition, parent: string, count: number): string {
    const link = parent + "-->"
    const style = count + "[\"" + node.name + "\"];\n"

    return link + style
}
