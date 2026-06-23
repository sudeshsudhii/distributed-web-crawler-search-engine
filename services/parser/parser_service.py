"""
Parser Service — HTML parsing, text extraction, and link extraction.

This service consumes raw HTML from the crawl pipeline and produces:
- Extracted text content (boilerplate removed)
- Discovered outgoing links
- Document metadata (title, language, word count)

Architecture: Standalone Python microservice communicating via Kafka.
"""

import logging

logging.basicConfig(
    level=logging.INFO,
    format="%(asctime)s [%(levelname)s] %(name)s: %(message)s",
)
logger = logging.getLogger("parser-service")


class HTMLExtractor:
    """Extracts main content from HTML, removing boilerplate (nav, footer, ads)."""

    def extract(self, html: str) -> dict:
        """Extract text, title, and metadata from HTML content."""
        # TODO: Implement using BeautifulSoup + readability algorithm
        raise NotImplementedError("HTML extraction not yet implemented")


class LinkExtractor:
    """Extracts and normalizes outgoing links from HTML."""

    def extract(self, html: str, base_url: str) -> list[str]:
        """Extract all href links, resolve relative URLs, normalize."""
        # TODO: Implement URL normalization and canonicalization
        raise NotImplementedError("Link extraction not yet implemented")


def main():
    logger.info("Parser Service starting...")
    logger.info("Status: scaffold — implementation pending")


if __name__ == "__main__":
    main()
