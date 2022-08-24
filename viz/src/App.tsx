import React from 'react';
import content from './data/data.json';
import DagVizualizer from "./components/dag/DagVizualizer";
import TopBar from "./components/ui/TopBar";
import {DagResults} from "./types";
import "./index.css";
import Tree from "./components/tree/Tree";

const parsed = content as DagResults;

function App() {
    return (
        <div>
            <TopBar path="temporary.cue"/>
            <div style={{ display: "flex", flexDirection: "row"}}>
                <div className={'halfScreen'}>
                    <DagVizualizer data={parsed}/>
                </div>
                <div className={'halfScreen'}>
                    <Tree data={parsed}/>
                </div>
            </div>
        </div>
    );
}

export default App;
