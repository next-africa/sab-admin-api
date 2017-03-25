/**
 * Created by pdiouf on 2017-03-24.
 */
import React from 'react'
const styles = {
    root :{
        fontFamily: '"HelveticaNeue-Light", "Helvetica Neue Light", "Helvetica Neue", Helvetica, Arial, "Lucida Grande", sans-serif',
        fontWeight: 300,
    },
    header: {
        backgroundColor: '#03a9f4',
        color: 'white',
        padding: '16px',
        fontSize: '1.5em',
    },
};

const ContentHeader = (props) => {
    const rootStyle = props.style ? {...styles.root, ...props.style} : styles.root;

    return (
        <div style={rootStyle}>
            <div style={styles.header}>{props.title}</div>
            {props.children}
        </div>
    );

};

ContentHeader.propTypes={
    styles: React.PropTypes.object,
    title: React.PropTypes.oneOfType([
        React.PropTypes.string,
        React.PropTypes.object,
    ]),
    children : React.PropTypes.object
}

export default ContentHeader;