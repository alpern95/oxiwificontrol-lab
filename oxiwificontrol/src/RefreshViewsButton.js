// in ./RefreshViewsButton.js
import * as React from "react";
import {
    Button,
    useUpdateMany,
    useRefresh,
    useNotify,
    Confirm,
    useUnselectAll,
} from 'react-admin';
import RefreshIcon from '@material-ui/icons/Refresh';

const RefreshViewsButton = ({ selectedIds }) => {
    const refresh = useRefresh();
    const notify = useNotify();
    const unselectAll = useUnselectAll();
    const [updateMany, { loading }] = useUpdateMany(
        'groupe/refresh',
        selectedIds,
        { views: 0 },
        {
            onSuccess: () => {
                refresh();
                notify('Borne updated');
                unselectAll('groupes');
            },
            onFailure: error => notify('Error: groupes not updated', 'warning'),
        }
    );

    return (
        <Button
            label="RefreshBorne"
            disabled={loading}
            onClick={updateMany}
        >
            <RefreshIcon />
        </Button>
    );
};

export default RefreshViewsButton;
