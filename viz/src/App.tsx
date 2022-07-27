import React from 'react';
import content from './data/data.json';
import DagVizualizer from "./components/DagVizualizer";

export type DagDefinition = {
  name: string,
  file: string,
  def: string,
  dependencies: string[],
}

function App() {
  return (
    <div>
      {/*JSON.stringify(content)*/}
      <DagVizualizer data={content}>
      </DagVizualizer>
    </div>
  );
}

export default App;
