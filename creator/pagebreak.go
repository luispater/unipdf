package creator

// PageBreak represents a page break for a chapter.
type PageBreak struct {
}

// newPageBreak create a new page break.
func newPageBreak() *PageBreak {
	return &PageBreak{}
}

// GeneratePageBlocks generates a page break block.
func (p *PageBreak) GeneratePageBlocks(ctx DrawContext) ([]*Block, DrawContext, error) {
	// Return two empty blocks.  First one simply means that there is nothing more to add at the current page.
	// The second one starts a new page.
	blocks := []*Block{
		NewBlock(ctx.PageWidth, ctx.PageHeight-ctx.Y),
		NewBlock(ctx.PageWidth, ctx.PageHeight),
	}

	// New Page. Place DrawContext in upper Left corner (with Margins).
	ctx.Page++
	newContext := ctx
	newContext.Y = ctx.Margins.Top
	newContext.X = ctx.Margins.Left
	newContext.Height = ctx.PageHeight - ctx.Margins.Top - ctx.Margins.Bottom
	newContext.Width = ctx.PageWidth - ctx.Margins.Left - ctx.Margins.Right
	ctx = newContext

	return blocks, ctx, nil
}
