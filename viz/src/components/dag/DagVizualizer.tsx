import React from 'react';
import Mermaid from "./Mermaid";
import {DagResults} from "../../types";
import generateChart from "../../parser";

interface DagVizualizerProps {
    data: DagResults,
}

const chart: string = `graph LR;
A[test]-->B;
B-->C;
B-->D[plop lanflz eknlzeknfz];
`

const chart2: string = `classDiagram
    Animal <|-- Duck
    Animal <|-- Fish
    Animal <|-- Zebra
    Animal : +int age
    Animal : +String gender
    Animal: +isMammal()
    Animal: +mate()
    class Duck{
        +String beakColor
        +swim()
        +quack()
    }
    class Fish{
        -int sizeInFeet
        -canEat()
    }
    class Zebra{
        +bool is_wild
        +run()
    }`

export default function DagVizualizer(props: DagVizualizerProps): JSX.Element {
    return (
        <Mermaid chart={generateChart(props.data)} />
    )
}
