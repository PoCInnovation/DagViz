import React from "react";
import MermaidReact from 'mermaid-react';

interface MermaidProps {
    chart: string;
}

export default function Mermaid(props: MermaidProps): JSX.Element {
    console.log(props.chart);
    return <MermaidReact
        id='test'
        mmd={props.chart}
        onClick={() => console.log('test Click')}
        onRender={svg => console.log('render content', svg)}
    />
}
