import React from 'react';
import content from './data/data.json';
import DagVizualizer from "./components/dag/DagVizualizer";
import TopBar from "./components/ui/TopBar";

export type DagDefinition = {
  name: string,
  file: string,
  def: string,
  dependencies: string[],
}

function App() {
  return (
    <div>
      <TopBar path="temporary.cue" />
      <DagVizualizer data={content} />
    </div>
  );
}

export default App;
