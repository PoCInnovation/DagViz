import React from 'react';
import content from './data/data.json';
import DagVizualizer from "./components/dag/DagVizualizer";
import TopBar from "./components/ui/TopBar";
import {DagResults} from "./types";

const parsed = content as DagResults;

function App() {
  return (
    <div>
      <TopBar path="temporary.cue" />
      <DagVizualizer data={parsed} />
    </div>
  );
}

export default App;
