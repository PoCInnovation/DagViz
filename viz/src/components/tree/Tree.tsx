// @ts-ignore
import FolderTree from "react-folder-tree";
import { Leaf } from "../../types";
import {Box} from "@mui/material";

type DagTree = {
  data: Leaf;
  onChange: (data: Leaf) => void;
};

export default function Tree({ data, onChange }: DagTree): JSX.Element {
  return (
      <Box marginTop={2} marginLeft={2}>
          <FolderTree
              data={data}
              showCheckbox={false}
              readOnly={true}
              initOpenStatus="custom"
              onChange={(state: any, event: any) => {
                  if (event.type !== "initialization") {
                      onChange(state);
                  }
              }}
          />
      </Box>
  );
}
