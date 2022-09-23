import { useState } from "react";
import DagVizualizer from "./components/dag/DagVizualizer";
import Tree from "./components/tree/Tree";
import TopBar from "./components/ui/TopBar";
import content from "./data/data.json";
import "./index.css";
import { generateTree } from "./parser";
import { DagResults, Leaf } from "./types";
import {baseColors, rootColor} from "./colors";

const parsed = content as DagResults;
const firstLeaf: Leaf = {
  name: "root",
  depth: 0,
  color: rootColor,
  colorParams: {
    shadeArray: baseColors,
    count: 0,
  },
  isOpen: false,
  checked: 0,
  metadata: {
    file: "root",
    def: "definition",
  },
}
const initialTree: Leaf[] = generateTree(parsed, firstLeaf);
const staticTree: Leaf = {
  ...firstLeaf,
  isOpen: true,
  children: initialTree
}


function App() {
  const [treeState, setTreeState] = useState(staticTree);

  return (
    <>
      <TopBar path={parsed.file} />
      <DagVizualizer file={content.file} data={treeState} />
      <Tree data={staticTree} onChange={setTreeState} />
    </>
  );
}

export default App;
