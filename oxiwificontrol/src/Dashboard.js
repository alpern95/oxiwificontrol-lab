// in src/Dashboard.js
import * as React from "react";
import { Card, CardContent, CardHeader } from '@material-ui/core';
import { useAuthenticated } from 'react-admin';

const MyPage = () => {
    useAuthenticated(); // redirects to login if not authenticated
    return (
        <div>
        <Card>
          <CardHeader title="Bienvenue sur OxiWifiControl" />
          <CardContent>Lorem ipsum sic dolor amet...</CardContent>
        </Card>
        </div>
    )
};

export default MyPage;
//export default () => (
//    <Card>
//        <CardHeader title="Welcome to the administration" />
//        <CardContent>Lorem ipsum sic dolor amet...</CardContent>
//    </Card>
//);
