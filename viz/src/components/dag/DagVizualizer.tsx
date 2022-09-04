import React from 'react';
import EChartsReact from "echarts-for-react";
import {Box} from "@mui/material";

interface DagVizualizerProps {
    file: string
}

export default function DagVizualizer(props: DagVizualizerProps): JSX.Element {
    const options = {
        grid: { top: 8, right: 8, bottom: 24, left: 36 },
        xAxis: {
            type: 'category',
            data: ['Mon', 'Tue', 'Wed', 'Thu', 'Fri', 'Sat', 'Sun'],
        },
        yAxis: {
            type: 'value',
        },
        series: [
            {
                data: [820, 932, 901, 934, 1290, 1330, 1320],
                type: 'line',
                smooth: true,
            },
        ],
        tooltip: {
            trigger: 'axis',
        },
    };

    return (
        <Box sx={{ border: 1, margin: 2 }}>
            <EChartsReact option={options} />
        </Box>
    );
}
