// in ./StopViewsButton.js
import * as React from "react";
import {
    Button,
    useUpdateMany,
    useRefresh,
    useNotify,
    useUnselectAll,
} from 'react-admin';
//import { VisibilityOff } from '@material-ui/icons';
import StopIcon from '@material-ui/icons/Stop';
//import PlayArrowIcon from '@material-ui/icons/PlayArrow';

const StopViewsButton = ({ selectedIds }) => {
    const refresh = useRefresh();
    const notify = useNotify();
    const unselectAll = useUnselectAll();
    const [updateMany, { loading }] = useUpdateMany(
        'groupe/stop',
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
            label="StopBorne"
            disabled={loading}
            onClick={updateMany}
        >
            <StopIcon />
        </Button>
    );
};

export default StopViewsButton;
