import React, {useEffect} from "react";
import mermaidAPI from "mermaid";

interface MermaidProps {
    chart: string;
}

mermaidAPI.initialize({
    startOnLoad: false,
});

const chart: string = `graph TD
        A[Client] --> B[Load Balancer]
        B --> C[Server01]
        B --> D[Server02]`

export default function Mermaid(props: MermaidProps): JSX.Element {
    useEffect(() => {
        mermaidAPI.initialize({
            maxTextSize: 1000000,
            startOnLoad: false,
            flowchart: {
                useMaxWidth: true,
                curve: 'cardinal',
            },
        })
        mermaidAPI.render('mermaid', props.chart, () => {

        });
    }, [])

    return <div className="mermaid">
        {chart}
    </div>
}
