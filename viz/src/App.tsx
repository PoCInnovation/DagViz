import { useState } from "react";
import DagVizualizer from "./components/dag/DagVizualizer";
import Tree from "./components/tree/Tree";
import TopBar from "./components/ui/TopBar";
import content from "./data/data.json";
import "./index.css";
import { generateTree } from "./parser";
import { DagResults, Leaf } from "./types";

const parsed = content as DagResults;
const initialTree: Leaf[] = generateTree(parsed);
const staticData = {
  name: "root",
  isOpen: true,
  checked: 0,
  children: initialTree,
  metadata: {
    file: "root",
    def: "salut",
  },
};

function App() {
  const [treeState, setTreeState] = useState(staticData);

  return (
    <>
      <TopBar path="temporary.cue" />
      <DagVizualizer file={content.file} flo={treeState} />
      <Tree data={staticData} onChange={setTreeState} />
    </>
  );
}

export default App;
