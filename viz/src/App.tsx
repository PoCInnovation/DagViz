import React, {useState} from 'react';
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
    const staticData = {
        name: 'root',
        isOpen: true,
        checked: 0,
        children: initialTree,
        metadata: {
            file: 'root',
            def: 'salut'
        }
    }

    const [treeState, setTreeState] = useState({
        name: 'root',
        isOpen: true,
        checked: 0,
        children: initialTree,
        metadata: {
            file: 'root',
            def: 'salut'
        }
    });

    return (
        <>
            <TopBar path="temporary.cue"/>
            <DagVizualizer file={content.file} flo={treeState}/>
            <Tree data={staticData} change={setTreeState}/>
        </>
    );
}

export default App;
