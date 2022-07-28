import React, {useEffect} from "react";
import mermaidAPI from "mermaid";
import MermaidReact from 'mermaid-react';

interface MermaidProps {
    chart: string;
}

mermaidAPI.initialize({
    startOnLoad: false,
});

const t = `
graph LR;
A-->B;
B-->C;
B-->D[plop lanflz eknlzeknfz];
`

export default function Mermaid(props: MermaidProps): JSX.Element {
    return <MermaidReact
        id='test'
        mmd={t}
        onClick={() => console.log('test Click')}
        onRender={svg => console.log('render content', svg)} />
}
