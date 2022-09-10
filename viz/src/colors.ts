import colorVariations from "color-variations";

export const rootColor = '#808080';

const baseColors = {
    blue: '#0000ff',
    green: '#00ff00',
    red: '#ff0000',
    yellow: '#ffff00',
    orange: '#ffa500',
    purple: '#800080',
    pink: '#ffb6c1',
    brown: '#a52a2a',
    aqua: '#00ffff',
    beige: '#ffffe0',
};

const baseColorsArray = Object.values(baseColors);
const baseColorMap = new Map<string, string>(baseColorsArray.map((c, i) => [c, Object.values(baseColors)[i]]));
let baseColorCount = 0;

const colors = colorVariations(baseColors, {
    steps: 10,
    includedFns: ['darken', 'desaturate', 'lighten', 'saturate', 'shade', 'tint'],
    excludedFns: [],
});

export function initColors() {
    baseColorCount = 0;
}

export function getColor(depth: number, parentColor: string): string {
    switch (depth) {
        case 1:
            return baseColorsArray[baseColorCount++];
        default:
            return parentColor
    }
}
