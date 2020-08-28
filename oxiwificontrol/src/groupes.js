import * as React from "react";
import { List, Datagrid, TextField } from 'react-admin';
// You can replace the list of default actions by your own element using the actions prop
import { Fragment } from 'react';
//import Button from '@material-ui/core/Button';
import StopViewsButton from './StopViewsButton';
import StartViewsButton from './StartViewsButton';
import RefreshViewsButton from './RefreshViewsButton';
//import { useQuery, Loading, Error } from 'react-admin';

//Add custom action

const PostBulkActionButtons = props => (
	    <Fragment>
            <StopViewsButton label="Stop Views" {...props} />
            <StartViewsButton label="Start Views" {...props} />
            <RefreshViewsButton label="Refresh Views" {...props} />
            {/* default bulk delete action */}
        </Fragment>
);

//Visu liste par groupe
export const GroupeList = ({  ...props }) => (
    <List {...props} title="Liste des bornes de votre groupe" bulkActionButtons={<PostBulkActionButtons />}>
        <Datagrid >
            <TextField source="nom" sortable={false} />
            <TextField source="description" sortable={false} /> 
            <TextField source="groupe" sortable={false} />
            <TextField source="etat" sortable={false} />
            <TextField source="lastrefresh" sortable={false} />
        </Datagrid>
    </List>
);
