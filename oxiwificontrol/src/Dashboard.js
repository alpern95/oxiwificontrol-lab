// in src/Dashboard.js
import * as React from "react";
import { Card, CardContent, CardHeader } from '@material-ui/core';
import { useAuthenticated } from 'react-admin';
import { makeStyles } from '@material-ui/core/styles';
//import { FieldProps, Product } from '../types';

const useStyles = makeStyles({
    root: { display: 'inline-block', marginTop: '1em', zIndex: 2 },
    content: { padding: 0, '&:last-child': { padding: 0 } },
    img: {
        width: 'initial',
        minWidth: 'initial',
        maxWidth: '42em',
        maxHeight: '15em',
    },
});

const MyPage = () => {
    const classes = useStyles();
    useAuthenticated(); // redirects to login if not authenticated
    return (
        <div>
        <Card>
          <CardHeader title="Bienvenue sur OxiWifiControl" />
            <CardContent className={classes.content}>
                <br />- Application de gestion des bornes wifi
                <br />
                <br />  Vous pourrez:
                <br />  - Allumer une borne wifi
                <br />  - Etteindre une borne wifi
                <br />  - Connaitre l'état d'une borne wifi

            </CardContent>
        </Card>

        </div>
    )
};

export default MyPage;
