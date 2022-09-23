import { Box } from "@mui/material";
import EChartsReact from "echarts-for-react";
import { generateChartInfo } from "../../parser";
import { Leaf } from "../../types";

interface DagVizualizerProps {
  file: string;
  data: Leaf;
}

export default function DagVizualizer(props: DagVizualizerProps): JSX.Element {
  const { data, links } = generateChartInfo(props.file, props.data);
  const options = {
    tooltip: [
      {
        show: true,
        showContent: true,
        trigger: "item",
        triggerOn: "mousemove",
        displayMode: "single",
        renderMode: "auto",
      },
    ],
    series: {
      type: "graph",
      id: "dagviz-graph",
      layout: "force",
      roam: true,
      emphasis: {
        disabled: false,
      },
      data,
      links,
    },
  };

  return (
      <EChartsReact option={options} />
  );
}
