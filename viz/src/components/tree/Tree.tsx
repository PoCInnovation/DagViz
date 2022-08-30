// @ts-ignore
import FolderTree from 'react-folder-tree';
import {DagResults} from "../../types";
import {generateTree} from "../../parser";

interface DagTree {
    data: DagResults,
}

export default function Tree(props: DagTree): JSX.Element {
    const treeState = {
        name: 'root',
        isOpen: false,
        children: generateTree(props.data),
        metadata: "hello world"
    }
    const treeState1 = {
        name: 'root [half checked and opened]',
        isOpen: false,   // this folder is opened, we can see it's children
        children: [
            {name: 'children 1 [not checked]', checked: 0},
            {
                name: 'children 2 [half checked and not opened]',
                checked: 0.5,
                isOpen: false,
                children: [
                    {name: 'children 2-1 [not checked]', checked: 0},
                    {name: 'children 2-2 [checked]', checked: 1},
                ],
            },
        ],
    };

    return (
        <FolderTree
            data={treeState}
            showCheckbox={ false }
            readOnly={ true }
            onNameClick={ (node: any) => {
                console.log(node)
            } }
            //onChange={ onTreeStateChange }
        />
    );
}
