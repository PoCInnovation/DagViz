import {DagDefinition, DagResults} from './types';

export default function generateChart(dag: DagResults): string {
    const style = "class diagram; direction LR;"
    const fileNode = "PARENT[\"" + dag.file + "\"]"
    let count = 0;

    const nodes = dag.dag.map(node => {
        count += 1
        return generateNode(node, fileNode, count);
    });

    return style + nodes.join("");
}

function generateNode(node: DagDefinition, parent: string, count: number): string {
    const link = parent + "-->"
    const style = count + "[\"" + node.name + "\"];"

    return link + style
}
