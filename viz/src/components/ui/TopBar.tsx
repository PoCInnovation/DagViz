import {AppBar, Box, Toolbar, Typography} from "@mui/material";

export interface TopBarProps {
    path: string
}

export default function TopBar(props: TopBarProps): JSX.Element {
    return (
        <Box sx={{ flexGrow: 1 }}>
            <AppBar position="static">
                <Toolbar>
                    <Typography variant="h6" component="div" sx={{ flexGrow: 1 }}>
                        DagViz
                    </Typography>
                    <Typography variant="subtitle1" component="div" sx={{ flexGrow: 1 }}>
                        File: {props.path}
                    </Typography>
                </Toolbar>
            </AppBar>
        </Box>
    )
}
