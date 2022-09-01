import React from 'react';
import Mermaid from "./Mermaid";
import {DagResults} from "../../types";
import {generateChart} from "../../parser";

interface DagVizualizerProps {
    data: string,
}

export default function DagVizualizer(props: DagVizualizerProps): JSX.Element {
    return (
        <Mermaid chart={props.data} />
    )
}
