-- noinspection SqlNoDataSourceInspectionForFile

-- First, insert the article sources
INSERT INTO note.article_source (source_id, name) VALUES
                                                      ('wired', 'Wired'),
                                                      ('the-verge', 'The Verge'),
                                                      (NULL, 'Gizmodo.com'),
                                                      (NULL, 'Android Central'),
                                                      (NULL, 'MacRumors'),
                                                      (NULL, 'CNET'),
                                                      ('business-insider', 'Business Insider'),
                                                      (NULL, 'Smashingmagazine.com'),
                                                      (NULL, 'Hackaday'),
                                                      (NULL, 'VentureBeat'),
                                                      ('the-next-web', 'The Next Web'),
                                                      ('time', 'Time'),
                                                      ('polygon', 'Polygon'),
                                                      (NULL, 'The New Yorker'),
                                                      (NULL, 'The Information'),
                                                      ('ign', 'IGN'),
                                                      ('hacker-news', 'Hacker News');

-- Insert a note for this search query
INSERT INTO note.note (status, total_results, query_text)
VALUES ('ok', 4205, 'technology news');

-- Insert articles
INSERT INTO note.article (
    source_id,
    author,
    title,
    description,
    url,
    url_to_image,
    published_at,
    content
) VALUES
-- Wired articles
(1, 'Lily Hay Newman', 'Stupid and Dangerous: CISA Funding Chaos Threatens Essential Cybersecurity Program',
 'The CVE Program is the primary way software vulnerabilities are tracked. Its long-term future remains in limbo even after a last-minute renewal of the US government contract that funds it.',
 'https://www.wired.com/story/cve-program-cisa-funding-chaos/',
 'https://media.wired.com/photos/67ffb7af15dd10f7d18e6cfb/191:100/w_1280,c_limit/CVE-nonprofit-sec-1301775422.jpg',
 '2025-04-16 20:10:04',
 'In an eleventh-hour scramble before a key contract was set to expire on Tuesday night, the United States Cybersecurity and Infrastructure Security Agency renewed its funding for the longtime software… [+3838 chars]'),

(1, 'Scharon Harding, Ars Technica', 'LG''s Integrated TV Ad Tech Analyzes Your Emotions',
 'LG has licensed tech that claims to interpret TV users'' feelings and convictions. The company will use this data to more directly target the ads it''s showing to users of its smart TV platform.',
 'https://www.wired.com/story/lg-zenapse-integrated-tv-ads-emotional-analysis/',
 '',
 '2025-04-18 10:30:00',
 'LG TVs will soon leverage an artificial intelligence model built for showing advertisements that more closely align with viewers'' personal beliefs and emotions.
 This story originally appeared on Ars… [+4180 chars]'),

-- The Verge articles
(2, 'Jess Weatherbed', 'Figma tells AI startup to stop using the term ''Dev Mode''',
 'Figma slapped Swedish AI coding startup Loveable with a cease-and-desist warning for naming one of its new product features "Dev Mode." It turns out Figma successfully trademarked the term Dev Mode in November last year, according to the US Patent and Tradema…',
 'https://www.theverge.com/news/649851/figma-dev-mode-trademark-loveable-dispute',
 'https://platform.theverge.com/wp-content/uploads/sites/2/chorus/uploads/chorus_asset/file/25515615/STK273_FIGMA.jpg?quality=90&strip=all&crop=0%2C10.732984293194%2C100%2C78.534031413613&w=1200',
 '2025-04-16 15:30:50',
 'Figma slapped Swedish AI coding startup Loveable with a cease-and-desist warning for naming one of its new product features "Dev Mode." It turns out Figma successfully trademarked the term Dev Mode i… [+1845 chars]'),

(2, 'Umar Shakir', 'Nintendo says the eShop will run more smoothly on the Switch 2',
 'If you''ve dreaded entering the eShop for most of the original Switch''s life because of how slow it is, well, Nintendo knows and won''t make that mistake again. The Switch 2 will have a faster-performing eShop channel "even when displaying a large number of gam…',
 'https://www.theverge.com/news/641966/nintendo-eshop-smooth-switch-2',
 'https://platform.theverge.com/wp-content/uploads/sites/2/2025/04/H2x1_NintendoeShop_WebsitePortal_enGB.jpg?quality=90&strip=all&crop=0,10.732984293194,100,78.534031413613',
 '2025-04-02 16:40:51',
 'The Nintendo eShop also has a new game discovery feature.
 The Nintendo eShop also has a new game discovery feature.
 If youve dreaded entering the eShop for most of the original Switchs life because… [+1367 chars]'),

(2, 'Jennifer Pattison Tuohy', 'Smart home device manufacturers are bracing for chaos — again',
 'President Donald Trumpâ\u0080\u0099s latest round of tariffs â\u0080\u0094 including a now 125 percent levy on Chinese imports â\u0080\u0094 will hit the smart home industry hard. Many smart home device makers are already struggling, thanks, in part, to Trumpâ\u0080\u0099s first round of tariffs. I…',
 'https://www.theverge.com/smart-home/645927/smart-home-device-manufacturers-are-bracing-for-chaos-again',
 'https://platform.theverge.com/wp-content/uploads/sites/2/2025/04/STKS488_TARIFFS_3_CVirginia_E.jpg?quality=90&strip=all&crop=0%2C10.732984293194%2C100%2C78.534031413613&w=1200',
 '2025-04-09 18:27:12',
 'New supply chain havoc caused by Trumps latest tariffs could derail the industrys growth.
 President Donald Trumps latest round of tariffs including a now 125 percent levy on Chinese imports will hit… [+8005 chars]'),

-- Gizmodo articles
(3, 'Kyle Barr', 'Nintendo is Bringing Us Kicking and Screaming Into the $80 Game Era with the Switch 2',
 'First-party Switch 2 games are now set at $80 for major releases, and other publishers will likely follow suit.',
 'https://gizmodo.com/nintendo-is-bringing-us-kicking-and-screaming-into-the-80-game-era-with-the-switch-2-2000584828',
 'https://gizmodo.com/app/uploads/2025/04/Nintendo-Switch-2-Hands-On-2.jpg',
 '2025-04-03 19:15:02',
 'Nintendo did not offer a price for its Switch 2 during its hour-long showcase. Instead, customers found out through the grapevine that the console would cost $450. But for many, the real eye-watering… [+6416 chars]'),

-- Android Central articles
(4, 'Nickolas Diaz', 'Xiaomi prepares developers for Android 16 with a preview on ''select'' devices',
 'Xiaomi announced the start of its Android 16 Developer Preview.',
 'https://www.androidcentral.com/phones/xiaomi/xiaomi-android-16-developer-preview-details-announced',
 'https://cdn.mos.cms.futurecdn.net/VJLNPqbRSbFtvBvFoPnAsG-1200-80.jpg',
 '2025-04-09 18:51:47',
 'What you need to know
 <ul><li>Xiaomi announces its Android 16 Developer Preview Program for those creating apps for its upcoming upgrade.</li><li>The preview will provide developers with the tools n… [+2739 chars]'),

-- CNET articles
(6, 'Jon Reed', 'Apple''s Folding iPhone Rumored to Have Face ID Embedded in Screen',
 'The latest leak suggests Apple has put its novel facial recognition technology into the long-rumored folding iPhone.',
 'https://www.cnet.com/tech/mobile/apples-folding-iphone-rumored-to-have-face-id-embedded-in-screen/',
 'https://www.cnet.com/a/img/resize/7c7d16f61d2df28cadde6f184c12a3403a8161df/hub/2025/04/14/e27c2407-b077-4558-9e93-b89f52739fa5/gettyimages-2209621854.jpg?auto=webp&fit=crop&height=675&width=1200',
 '2025-04-14 21:04:00',
 'As anticipation builds for a potential foldable iPhone, one rumor suggests it may have another feature that has yet to appear in another Apple phone: a camera embedded under the screen for Face ID.
 … [+1628 chars]'),

-- Business Insider articles
(7, 'Rebecca Torrence', 'Meet the two twenty-somethings running Palantir''s healthcare AI business',
 'Palantir says it''s helping top health systems save millions of dollars with its custom AI-powered software solutions.',
 'https://www.businessinsider.com/meet-two-twenty-somethings-running-palantir-healthcare-business-2025-4',
 'https://i.insider.com/67ec35f5b8b41a9673fc523a?width=1200&format=jpeg',
 '2025-04-03 09:00:03',
 'Jeremy David and Drew Goldstein are co-heads of healthcare at Palantir.Palantir
 <ul><li>Palantir launched its healthcare business four years ago, with two 25-year-olds at the helm.</li><li>The $193 … [+5897 chars]'),

-- Time articles
(12, 'Billy Perrigo/Austin, Texas', 'Inside Amazon''s Race to Build the AI Industry''s Biggest Datacenters',
 'Amazon and Anthropic are racing to build the world''s largest AI datacenters, to challenge OpenAI and Microsoft''s $100 billion ''Stargate.''',
 'https://time.com/7273288/amazon-anthropic-openai-microsoft-stargate-datacenters/',
 'https://api.time.com/wp-content/uploads/2025/04/AWS-Trainium.jpg?quality=85&w=1200&h=628&crop=1',
 '2025-04-02 14:56:12',
 'Rami Sinno is crouched beside a filing cabinet, wrestling a beach-ball sized disc out of a box, when a dull thump echoes around his laboratory.
 I just dropped tens of thousands of dollars worth of m… [+9325 chars]');

-- Associate articles with the note
INSERT INTO note.note_article (note_id, article_id)
SELECT 1, id FROM note.article;