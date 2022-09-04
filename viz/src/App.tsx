import React from 'react';
import content from './data/data.json';
import DagVizualizer from "./components/dag/DagVizualizer";
import {DagResults, Leaf} from "./types";
import "./index.css";
import {generateChartInfo, generateTree} from "./parser";
import TopBar from "./components/ui/TopBar";
import Tree from "./components/tree/Tree";
import {Box} from "@mui/material";

const parsed = content as DagResults;
const initialTree: Leaf[] = generateTree(parsed);

function App() {
    const { data, links } = generateChartInfo(initialTree);

    return (
        <>
            <TopBar path="temporary.cue"/>
            <DagVizualizer file={content.file} data={data} links={links} />
        </>
    );
}

export default App;
