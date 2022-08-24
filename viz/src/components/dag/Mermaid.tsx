import React, {useEffect} from "react";
import mermaid from "mermaid";
import { MERMAID_CONFIG } from "./mermaidConfig";

interface MermaidProps {
    chart: string
}

export default function Mermaid(props: MermaidProps): JSX.Element {
    mermaid.initialize(MERMAID_CONFIG)

    useEffect(() => {
        mermaid.contentLoaded()
    }, [props.chart])

    return (
        <div className="mermaid">
            {props.chart}
        </div>
    )
}
