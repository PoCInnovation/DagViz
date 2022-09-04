import React from 'react';
import content from './data/data.json';
import DagVizualizer from "./components/dag/DagVizualizer";
import {DagResults, Leaf} from "./types";
import "./index.css";
import {generateChartInfo, generateTree} from "./parser";

const parsed = content as DagResults;
const initialTree: Leaf[] = generateTree(parsed);

function App() {
    const [tree, setTree] = React.useState(initialTree);
    const { data, links } = generateChartInfo(initialTree);

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
        <DagVizualizer file={content.file} data={data} links={links} />
    );
}

export default App;
