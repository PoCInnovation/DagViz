import React from 'react';
import EChartsReact from "echarts-for-react";
import {Box} from "@mui/material";

interface DagVizualizerProps {
    file: string
    data: any,
    links: any,
}

export default function DagVizualizer(props: DagVizualizerProps): JSX.Element {
    console.log(props.links)

    const options = {
        series: {
            type: "graph",
            id: "dagviz-graph",
            layout: "force",
            roam: true,
            emphasis: {
                disabled: false,
            },
            data: props.data,
            links: props.links,
        }
    };

    return (
        <Box sx={{ border: 1, margin: 2 }}>
            <EChartsReact option={options} />
        </Box>
    );
}
