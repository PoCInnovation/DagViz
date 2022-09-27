import colorVariations from "color-variations";
import {Leaf} from "./types";

export const rootColor = '#808080';
export const maxDepth = 3;

export const baseColors: string[] = [
    '#0000ff',
    '#00ff00',
    '#ffff00',
    '#ff0000',
    '#ffa500',
    '#800080',
    '#ffb6c1',
    '#a52a2a',
    '#00ffff',
];

type ColorResults = {
    colorShade: string[]
}

export function generateColors(currentLeaf: Leaf, parentLeaf: Leaf) {
    if (currentLeaf.depth <= maxDepth) {
        const colorGen: string = parentLeaf.colorParams.shadeArray[parentLeaf.colorParams.count];
        const variations = colorVariations( { color: colorGen }, { steps: 8, includedFns: ['shade'], excludedFns: [] })

        currentLeaf.color = colorGen;
        currentLeaf.colorParams = {
            shadeArray: (variations as ColorResults).colorShade,
            count: 0
        }
        parentLeaf.colorParams.count = parentLeaf.colorParams.count === parentLeaf.colorParams.shadeArray.length - 1 ? 0 : parentLeaf.colorParams.count + 1;
    } else {
        currentLeaf.color = parentLeaf.color
    }

    /*
    switch (depth) {
        case 1:
            return baseColorsArray[baseColorCount++]
        case 2:
            const baseColor = baseColorMap.get(parentColor) || "blue";
            return colors[baseColor + "shade"][0];
        default:
            return parentColor
    }
    */
}
