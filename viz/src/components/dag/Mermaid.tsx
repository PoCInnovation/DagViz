import React from "react";
import MermaidReact from 'mermaid-react';

interface MermaidProps {
    chart: string;
}

function callback() {
    console.log('bonjour');
}

export default function Mermaid(props: MermaidProps): JSX.Element {
    const a = "graph LR\n" +
        "ABC(Datasource) -- Input --> B((System))\n" +
        "B -- Output --> C(Datasink)\n" +
        "click ABC callback"

    return <MermaidReact
        id='test'
        mmd={a}
        onClick={() => console.log('test Click')}
        onRender={svg => console.log('render content', svg)}
    />
}
