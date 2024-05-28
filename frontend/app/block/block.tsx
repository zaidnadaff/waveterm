// Copyright 2024, Command Line Inc.
// SPDX-License-Identifier: Apache-2.0

import { PlotView } from "@/app/view/plotview";
import { PreviewView } from "@/app/view/preview";
import { TerminalView } from "@/app/view/term";
import { CenteredDiv } from "@/element/quickelems";
import * as WOS from "@/store/wos";
import * as React from "react";

import "./block.less";

const Block = ({ tabId, blockId }: { tabId: string; blockId: string }) => {
    const blockRef = React.useRef<HTMLDivElement>(null);
    const [dims, setDims] = React.useState({ width: 0, height: 0 });

    function handleClose() {
        WOS.DeleteBlock(blockId);
    }

    React.useEffect(() => {
        if (!blockRef.current) {
            return;
        }
        const rect = blockRef.current.getBoundingClientRect();
        const newWidth = Math.floor(rect.width);
        const newHeight = Math.floor(rect.height);
        if (newWidth !== dims.width || newHeight !== dims.height) {
            setDims({ width: newWidth, height: newHeight });
        }
    }, [blockRef.current]);

    let blockElem: JSX.Element = null;
    const [blockData, blockDataLoading] = WOS.useWaveObjectValue<Block>(WOS.makeORef("block", blockId));
    if (blockDataLoading) {
        blockElem = <CenteredDiv>Loading...</CenteredDiv>;
    } else if (blockData.view === "term") {
        blockElem = <TerminalView blockId={blockId} />;
    } else if (blockData.view === "preview") {
        blockElem = <PreviewView blockId={blockId} />;
    } else if (blockData.view === "plot") {
        blockElem = <PlotView />;
    }
    return (
        <div className="block" ref={blockRef}>
            <div key="header" className="block-header">
                <div className="block-header-text text-fixed">
                    Block [{blockId.substring(0, 8)}] {dims.width}x{dims.height}
                </div>
                <div className="flex-spacer" />
                <div className="close-button" onClick={() => handleClose()}>
                    <i className="fa fa-solid fa-xmark-large" />
                </div>
            </div>
            <div key="content" className="block-content">
                <React.Suspense fallback={<CenteredDiv>Loading...</CenteredDiv>}>{blockElem}</React.Suspense>
            </div>
        </div>
    );
};

export { Block };
