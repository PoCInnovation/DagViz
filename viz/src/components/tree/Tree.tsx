// @ts-ignore
import FolderTree from "react-folder-tree";
import { Leaf } from "../../types";

type DagTree = {
  data: Leaf;
  onChange: (data: Leaf) => void;
};

export default function Tree({ data, onChange }: DagTree): JSX.Element {
  return (
    <FolderTree
      data={data}
      showCheckbox={false}
      readOnly={true}
      onNameClick={(node: any) => {
        console.log(node);
      }}
      initOpenStatus="custom"
      onChange={(state: any, event: any) => {
        if (event.type !== "initialization") {
          onChange(state);
        }
      }}
    />
  );
}
