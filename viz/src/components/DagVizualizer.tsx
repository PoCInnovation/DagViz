import React from 'react';
import {DagDefinition} from "../App";
import Mermaid from "./Mermaid";

export interface DagVizualizerProps {
    data: DagDefinition[];
    children: React.ReactNode
}

const chart: string = `graph LR;
A-->B;
B-->C;
B-->D[plop lanflz eknlzeknfz];
      `

export default function DagVizualizer(data: DagVizualizerProps): JSX.Element {

    return (
        <Mermaid chart={chart} />
    )
}
