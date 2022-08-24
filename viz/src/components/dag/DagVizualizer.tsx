import React from 'react';
import Mermaid from "./Mermaid";
import {DagResults} from "../../types";
import {generateChart} from "../../parser";

interface DagVizualizerProps {
    data: DagResults,
}

export default function DagVizualizer(props: DagVizualizerProps): JSX.Element {
    return (
        <Mermaid chart={generateChart(props.data)} />
    )
}
