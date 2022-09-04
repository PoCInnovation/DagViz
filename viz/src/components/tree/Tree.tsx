// @ts-ignore
import FolderTree from 'react-folder-tree';
import {Leaf} from "../../types";

interface DagTree {
    data: Leaf,
    change: (data: Leaf) => void
}

export default function Tree(props: DagTree): JSX.Element {

    /*const treeState1 = {
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
    };*/

    return (
        <FolderTree
            data={props.data}
            showCheckbox={ false }
            readOnly={ true }
            onNameClick={ (node: any) => {
                console.log(node)
            } }
            initOpenStatus='custom'
            onChange={ (state: any, event: any) => {
                    console.log(state, event)
                    props.change(state)
            } }
        />
    );
}
