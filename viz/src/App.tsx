import React, {useEffect} from 'react';
import content from './data/data.json';
import DagVizualizer from "./components/dag/DagVizualizer";
import TopBar from "./components/ui/TopBar";
import {DagResults, Leaf} from "./types";
import "./index.css";
import Tree from "./components/tree/Tree";
import {generateChart, generateTree} from "./parser";

const parsed = content as DagResults;
const initialTree: Leaf[] = generateTree(parsed);

function App() {
    const [tree, setTree] = React.useState(initialTree);
    const [displayTree, setDisplayTree] = React.useState("");

    function updateDisplay(current: Leaf[]): void {
        setDisplayTree(generateChart(current));
    }

    useEffect(() => {
        updateDisplay(initialTree);
    });

    return (
        /*
        <div>
            <TopBar path="temporary.cue"/>
            <div style={{ display: "flex", flexDirection: "row"}}>
                <div className={'halfScreen'}>
                    <DagVizualizer data={displayTree}/>
                </div>
                <div className={'halfScreen'}>
                    <Tree data={initialTree}/>
                </div>
            </div>
        </div>
        */
        <DagVizualizer file={content.file}/>
    );
}

export default App;
